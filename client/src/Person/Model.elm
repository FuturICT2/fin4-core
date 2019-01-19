module Person.Model exposing (Model, init)

import Array exposing (Array)
import Http
import Material
import Model.Persons exposing (Person)
import Person.Msg exposing (Msg(..))


type alias Model =
    { mdl : Material.Model
    , data : Maybe Person
    , tab : Int
    , error : Maybe String
    }


init : Model
init =
    { mdl = Material.model
    , data = Nothing
    , tab = 0
    , error = Nothing
    }
