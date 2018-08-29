module Portfolio.Command exposing (..)

import Main.Context exposing (Context)
import Portfolio.Msg exposing (Msg(..))
import Main.Routing exposing (Route(..))
import Common.Api exposing (get)
import Json.Decode as JD
import Json.Encode as JE
import Common.Json exposing (decodeAt, deocdeIntWithDefault)
import Model.Portfolio exposing (portfolioDecoder)


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
        "/portfolio/positions"
        []
        portfolioDecoder
