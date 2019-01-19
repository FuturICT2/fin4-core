module Tokens.Command exposing
    ( commands
    , likeCmd
    , loadTokensCmd
    )

import Common.Api exposing (get, postWithCsrf)
import Common.Json exposing (decodeAt, deocdeIntWithDefault, emptyResponseDecoder)
import Dict
import Json.Decode as JD
import Json.Encode as JE
import Main.Context exposing (Context)
import Main.Routing exposing (Route(..))
import Model.Tokens exposing (tokensDecoder)
import Tokens.Model exposing (Model)
import Tokens.Msg exposing (Msg(..))


commands : Context -> Cmd Msg
commands ctx =
    case ctx.route of
        TokensRoute ->
            loadTokensCmd ctx 1

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


loadTokensCmd ctx page =
    get ctx
        OnLoadTokensResponse
        "/tokens"
        []
        tokensDecoder
