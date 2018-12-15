module Actions.Command exposing (addRewardsCmd, commands, loadActionsCmd)

import Actions.Model exposing (Model)
import Actions.Msg exposing (Msg(..))
import Common.Api exposing (get, postWithCsrf)
import Common.Json exposing (decodeAt, deocdeIntWithDefault, emptyResponseDecoder)
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


addRewardsCmd : Context -> Model -> Int -> String -> Cmd Msg
addRewardsCmd ctx model actionId amount =
    postWithCsrf ctx
        OnAddRewardsResponse
        "/add-rewards-to-action"
        (encodeAddRewardsAction model actionId amount)
        emptyResponseDecoder


encodeAddRewardsAction : Model -> Int -> String -> JE.Value
encodeAddRewardsAction model actionId amount =
    JE.object
        [ ( "actionId", JE.int actionId )
        , ( "tokenId", JE.int 1 )
        , ( "amount", JE.string amount )
        ]


loadActionsCmd ctx page =
    get ctx
        OnLoadActionsResponse
        "/actions"
        []
        actionsDecoder
