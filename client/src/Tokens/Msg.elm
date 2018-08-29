module Tokens.Msg exposing (Msg(..))

import Http
import Material
import Model.Tokens exposing (Tokens)


type Msg
    = Mdl (Material.Msg Msg)
    | OnLoadTokensResponse (Result Http.Error Tokens)
