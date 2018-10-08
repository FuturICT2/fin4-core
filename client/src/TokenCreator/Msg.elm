module TokenCreator.Msg exposing (Msg(..))

import Http
import Material


type Msg
    = Mdl (Material.Msg Msg)
    | SetName String
    | SetTotalSupply String
