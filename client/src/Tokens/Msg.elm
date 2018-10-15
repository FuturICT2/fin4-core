module Tokens.Msg exposing (Msg(..))

import Http
import Main.User exposing (Users)
import Material
import Model.Tokens exposing (Tokens)


type Msg
    = Mdl (Material.Msg Msg)
    | OnLoadTokensResponse (Result Http.Error Tokens)
    | SelectTab Int
    | TickerTimout
