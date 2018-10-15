module Common.Api exposing (CSRF, addQueryToUrl, csrfCmd, deleteWithCsrf, get, postWithCsrf)

import Http
import HttpBuilder exposing (..)
import Json.Decode as JD
import Json.Decode.Pipeline as JP
import Main.Context exposing (Context)
import Task exposing (Task)


type alias CSRF =
    { token : String
    }


{-| takes a uri and a list of query tuples and return combined URI
takes:
"/api/user" [("id", "1"), ("foo","bar")]
returns:
"/api/user?id=1&foo=bar"
-}
addQueryToUrl : String -> List ( String, String ) -> String
addQueryToUrl uri args =
    let
        internal_ uri args =
            case args of
                [] ->
                    uri

                h :: t ->
                    let
                        ( n, v ) =
                            h
                    in
                    internal_ (uri ++ n ++ "=" ++ v ++ "&") t
    in
    internal_ (uri ++ "?") args


{-| creates a HTTP get request
ctx: application main context
msg: Msg to emit when http request is done
uri: URI to call "/api/user"
args: URI query params [("id", "1")]
decoder: json decoder of the expected result
-}
get ctx msg uri args decoder =
    Task.attempt msg <|
        Http.toTask
            (Http.get
                (ctx.flags.apiBase ++ addQueryToUrl uri args)
                decoder
            )


{-| csrfCmd a command to get CSRF token from serv
-}
csrfCmd : Context -> Http.Request CSRF
csrfCmd ctx =
    Http.get
        (ctx.flags.apiBase ++ "/csrf")
        (JP.decode CSRF |> JP.required "token" JD.string)


{-| gets a CSRF token from server and creates a HTTP post task
-}
postWithCsrf ctx msg uri jsonBody decoder =
    Task.attempt msg
        (Http.toTask
            (HttpBuilder.post (ctx.flags.apiBase ++ uri)
                |> withJsonBody jsonBody
                |> withExpect (Http.expectJson decoder)
                |> toRequest
            )
        )


deleteWithCsrf ctx msg uri jsonBody decoder =
    Task.attempt msg
        (Http.toTask (csrfCmd ctx)
            |> Task.andThen
                (\{ token } ->
                    Http.toTask
                        (HttpBuilder.delete (ctx.flags.apiBase ++ uri)
                            |> withJsonBody jsonBody
                            |> withExpect (Http.expectJson decoder)
                            |> withHeader "X-Csrf-Token" token
                            |> toRequest
                        )
                )
        )
