module Timeline.Command exposing
    ( commands
    , loadTimelineCmd
    , toggleFavoriteCmd
    , verifyBlockCmd
    )

import Common.Api exposing (get, postWithCsrf)
import Common.Json exposing (emptyResponseDecoder)
import Json.Decode.Pipeline as JP
import Json.Encode as JE
import Main.Context exposing (Context)
import Main.Routing exposing (Route(..))
import Timeline.Model exposing (TimelineType(..), timelineEntriesDecoder)
import Timeline.Msg exposing (Msg(..))


commands : Context -> TimelineType -> Cmd Msg
commands ctx timelineType =
    case ctx.route of
        HomepageRoute ->
            Cmd.batch
                [ loadTimelineCmd ctx 1 timelineType
                ]

        AssetRoute assetId ->
            Cmd.batch
                [ loadTimelineCmd ctx 1 timelineType
                ]

        ProfileRoute userId ->
            Cmd.batch
                [ loadTimelineCmd ctx 1 timelineType
                ]

        _ ->
            Cmd.none


toggleFavoriteCmd : Context -> Int -> Cmd Msg
toggleFavoriteCmd ctx blockId =
    postWithCsrf ctx
        OnToggleFavoriteResponse
        ("/asset-block/" ++ toString blockId ++ "/toggle-favorite")
        (JE.object [])
        (JP.decode {})


loadTimelineCmd : Context -> Int -> TimelineType -> Cmd Msg
loadTimelineCmd ctx page timelineType =
    let
        apiPath =
            case timelineType of
                HomepageTimeline ->
                    "/timeline"

                UserTimeline userId ->
                    "/timeline/user/" ++ toString userId

                AssetTimeline assetId ->
                    "/timeline/asset/" ++ toString assetId
    in
    get ctx
        OnLoadTimelineResponse
        apiPath
        [ ( "page", toString page ) ]
        timelineEntriesDecoder


verifyBlockCmd : Context -> Bool -> Int -> Cmd Msg
verifyBlockCmd ctx isAccepted blockId =
    postWithCsrf ctx
        VerifyBlockResponse
        ("/asset-block/" ++ toString blockId ++ "/verify")
        (JE.object
            [ ( "isAccepted", JE.bool isAccepted )
            ]
        )
        emptyResponseDecoder
