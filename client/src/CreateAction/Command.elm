module CreateAction.Command exposing (createActionCmd)

import Common.Api exposing (postWithCsrf)
import CreateAction.Model exposing (Model)
import CreateAction.Msg exposing (Msg(..))
import Http
import HttpBuilder exposing (..)
import Json.Decode as Decode
import Json.Decode.Pipeline as JP
import Json.Encode as JE
import Main.Context exposing (Context)
import Model.Tokens exposing (tokenDecoder)


createActionCmd : Context -> Model -> Cmd Msg
createActionCmd ctx model =
    postWithCsrf ctx
        OnCreateActionResponse
        "/actions"
        (encodeCreateAction model)
        tokenDecoder


encodeCreateAction : Model -> JE.Value
encodeCreateAction model =
    JE.object
        [ ( "description", JE.string model.description )
        , ( "timeLimit", JE.string "0" )
        ]
