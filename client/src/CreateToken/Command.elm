module CreateToken.Command exposing (createTokenCmd)

import Common.Api exposing (postWithCsrf)
import CreateToken.Model exposing (Model)
import CreateToken.Msg exposing (Msg(..))
import Http
import HttpBuilder exposing (..)
import Json.Decode as Decode
import Json.Decode.Pipeline as JP
import Json.Encode as JE
import Main.Context exposing (Context)
import Model.Tokens exposing (tokenDecoder)


createTokenCmd : Context -> Model -> Cmd Msg
createTokenCmd ctx model =
    postWithCsrf ctx
        OnCreateTokenResponse
        "/create-token"
        (encodeCreateToken model)
        tokenDecoder


encodeCreateToken : Model -> JE.Value
encodeCreateToken model =
    JE.object
        [ ( "name", JE.string model.name )
        , ( "purpose", JE.string model.description )
        , ( "totalSupply", JE.string "0" )
        , ( "symbol", JE.string model.symbol )
        ]
