module Portfolio.Command exposing (commands, loadPortfolioCmd)

import Common.Api exposing (get)
import Common.Json exposing (decodeAt, deocdeIntWithDefault)
import Json.Decode as JD
import Json.Encode as JE
import Main.Context exposing (Context)
import Main.Routing exposing (Route(..))
import Model.Portfolio exposing (portfolioDecoder)
import Portfolio.Msg exposing (Msg(..))


commands : Context -> Cmd Msg
commands ctx =
    case ctx.route of
        PortfolioRoute ->
            loadPortfolioCmd ctx 1

        _ ->
            Cmd.none


loadPortfolioCmd ctx page =
    get ctx
        OnLoadPortfolioResponse
        "/balances"
        []
        portfolioDecoder
