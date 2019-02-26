module CreateAsset.Command exposing (createAssetCmd)

import Asset.Model exposing (assetDecoder)
import Common.Api exposing (postWithCsrf)
import CreateAsset.Model exposing (Model)
import CreateAsset.Msg exposing (Msg(..))
import Json.Encode as JE
import Main.Context exposing (Context)


createAssetCmd : Context -> Model -> Cmd Msg
createAssetCmd ctx model =
    postWithCsrf ctx
        OnCreateAssetSuccess
        "/assets"
        (encodeCreateAsset model)
        assetDecoder


encodeCreateAsset : Model -> JE.Value
encodeCreateAsset model =
    JE.object
        [ ( "name", JE.string model.name )
        , ( "symbol", JE.string model.symbol )
        , ( "purpose", JE.string model.purpose )
        , ( "isSensor", JE.bool model.isSensor )
        ]
