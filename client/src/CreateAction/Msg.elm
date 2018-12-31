module CreateAction.Msg exposing (Msg(..))

import Http
import Material
import Model.Tokens exposing (Token)


type Msg
    = Mdl (Material.Msg Msg)
    | StepBack
    | StepForward
    | PostAction
    | OnCreateActionResponse (Result Http.Error Token)
    | SetTimeLimit String
    | SetDescription String
