module Actions.Command exposing
    ( commands
    , likeCmd
    , loadActionsCmd
    )

import Actions.Model exposing (Model)
import Actions.Msg exposing (Msg(..))
import Common.Api exposing (get, postWithCsrf)
import Common.Json exposing (decodeAt, deocdeIntWithDefault, emptyResponseDecoder)
import Dict
import Json.Decode as JD
import Json.Encode as JE
import Main.Context exposing (Context)
import Main.Routing exposing (Route(..))
import Model.Actions exposing (actionsDecoder)


commands : Context -> Cmd Msg
commands ctx =
    case ctx.route of
        ActionsRoute ->
            loadActionsCmd ctx 1

        _ ->
            Cmd.none


likeCmd ctx tokenId state =
    postWithCsrf ctx
        OnDoLikeResponse
        "/toggle-token-like"
        (JE.object
            [ ( "tokenId", JE.int tokenId )
            ]
        )
        emptyResponseDecoder


loadActionsCmd ctx page =
    get ctx
        OnLoadActionsResponse
        "/tokens"
        []
        actionsDecoder
