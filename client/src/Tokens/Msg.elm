module Tokens.Msg exposing (Msg(..))

import Common.Json exposing (EmptyResponse)
import Http
import Material
import Model.Tokens exposing (Tokens)


type Msg
    = Mdl (Material.Msg Msg)
    | OnLoadTokensResponse (Result Http.Error Tokens)
    | TickerTimout
    | OnDoLikeResponse (Result Http.Error EmptyResponse)
    | DoLike Int Bool
