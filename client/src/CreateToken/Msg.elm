module CreateToken.Msg exposing (Msg(..))

import Http
import Material
import Model.Tokens exposing (Token)


type Msg
    = Mdl (Material.Msg Msg)
    | StepBack
    | StepForward
    | PostToken
    | OnCreateTokenSuccess (Result Http.Error Token)
    | SetName String
    | SetDescription String
    | SetNewHashtag String
    | SetSymbol String
    | AddHashtag
    | RemoveHashtag Int
    | SetShares String
    | SetActiveCategory String
