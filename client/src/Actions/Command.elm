module Actions.Command exposing (commands, loadActionsCmd)

import Actions.Msg exposing (Msg(..))
import Common.Api exposing (get)
import Common.Json exposing (decodeAt, deocdeIntWithDefault)
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


loadActionsCmd ctx page =
    get ctx
        OnLoadActionsResponse
        "/actions"
        []
        actionsDecoder
