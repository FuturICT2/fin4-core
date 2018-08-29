module Tokens.Command exposing (commands, loadTokensCmd)

import Common.Api exposing (get)
import Common.Json exposing (decodeAt, deocdeIntWithDefault)
import Json.Decode as JD
import Json.Encode as JE
import Main.Context exposing (Context)
import Main.Routing exposing (Route(..))
import Model.Tokens exposing (tokensDecoder)
import Tokens.Msg exposing (Msg(..))


commands : Context -> Cmd Msg
commands ctx =
    case ctx.route of
        TokensRoute ->
            loadTokensCmd ctx 1

        _ ->
            Cmd.none


loadTokensCmd ctx page =
    get ctx
        OnLoadTokensResponse
        "/tokens"
        []
        tokensDecoder
