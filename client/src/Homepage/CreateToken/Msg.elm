module CreateToken.Msg exposing (Msg(..))

import Http
import Material
import Model.Tokens exposing (Token)


type Msg
    = Mdl (Material.Msg Msg)
    | StepBack
    | StepForward
    | PostToken
    | OnCreateTokenResponse (Result Http.Error Token)
    | SetName String
    | SetDescription String
    | SetSymbol String
    | SetShares String
