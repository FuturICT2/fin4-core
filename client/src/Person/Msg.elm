module Person.Msg exposing (Msg(..))

import Common.Json exposing (EmptyResponse)
import Http
import Material
import Model.Persons exposing (Person)


type Msg
    = Mdl (Material.Msg Msg)
    | OnPersonLoadResponse (Result Http.Error Person)
    | SelectTab Int
