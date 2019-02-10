module Asset.Command exposing
    ( commands
    , encodeBlockSubmition
    , encodeImage
    , loadAssetCmd
    , submitBlockCmd
    , verifyBlockCmd
    )

import Asset.Model exposing (Image, assetDecoder)
import Asset.Msg exposing (Msg(..))
import Common.Api exposing (get, postWithCsrf)
import Common.Json
    exposing
        ( decodeAt
        , deocdeIntWithDefault
        , emptyResponseDecoder
        )
import Json.Encode as JE
import Main.Context exposing (Context)
import Main.Routing exposing (Route(..))
import Timeline.Command
import Timeline.Model exposing (TimelineType(..))


commands : Context -> Cmd Msg
commands ctx =
    case ctx.route of
        AssetRoute assetId ->
            Cmd.batch
                [ loadAssetCmd ctx assetId
                ]

        _ ->
            Cmd.none


loadAssetCmd : Context -> Int -> Cmd Msg
loadAssetCmd ctx assetId =
    get ctx
        OnAssetLoadResponse
        ("/v2/assets/" ++ toString assetId)
        []
        assetDecoder


submitBlockCmd : Context -> Int -> String -> List Image -> Cmd Msg
submitBlockCmd ctx assetId blockText imgs =
    postWithCsrf ctx
        SubmitBlockResponse
        "/v2/create-asset-block"
        (encodeBlockSubmition assetId blockText imgs)
        emptyResponseDecoder


verifyBlockCmd : Context -> Bool -> Int -> Cmd Msg
verifyBlockCmd ctx isAccepted blockId =
    postWithCsrf ctx
        VerifyBlockResponse
        "/v2/verify-asset-block"
        (JE.object
            [ ( "blockId", JE.int blockId )
            , ( "isAccepted", JE.bool isAccepted )
            ]
        )
        emptyResponseDecoder


encodeBlockSubmition : Int -> String -> List Image -> JE.Value
encodeBlockSubmition assetId blockText images =
    JE.object
        [ ( "assetId", JE.int assetId )
        , ( "blockText", JE.string blockText )
        , ( "images", JE.list <| List.map encodeImage images )
        ]


encodeImage : Image -> JE.Value
encodeImage v =
    JE.object
        [ ( "contents", JE.string v.contents )
        , ( "filename", JE.string v.filename )
        ]
