module UserLogin.Command exposing
    ( encodeLogin
    , encodeSignupPayload
    , loginCmd
    , signupCmd
    , withDefault
    )

import Common.Api exposing (postWithCsrf)
import Http
import HttpBuilder exposing (..)
import Json.Decode as Decode
import Json.Decode.Pipeline as JP
import Json.Encode as JE
import Main.Context exposing (Context)
import Main.User exposing (User, userDecoder)
import RemoteData exposing (RemoteData)
import Task exposing (Task)
import UserLogin.Model exposing (Model)
import UserLogin.Msg exposing (Msg(..))


signupCmd : Context -> Model -> Cmd Msg
signupCmd ctx model =
    postWithCsrf ctx
        OnSignupResponse
        "/register"
        (encodeSignupPayload model)
        userDecoder


encodeSignupPayload : Model -> JE.Value
encodeSignupPayload model =
    JE.object
        [ ( "name", JE.string model.name )
        ]


loginCmd : Context -> Model -> Cmd Msg
loginCmd ctx model =
    postWithCsrf ctx
        OnLoginResponse
        "/login"
        (encodeLogin model)
        userDecoder


encodeLogin : Model -> JE.Value
encodeLogin model =
    JE.object
        [ ( "name", JE.string model.name )
        ]


withDefault : a -> Maybe a -> a
withDefault default maybeVal =
    case maybeVal of
        Just v ->
            v

        Nothing ->
            default
