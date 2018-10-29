module Actions.Model exposing (Model, init)

import Material
import Model.Actions exposing (Actions)


type alias Model =
    { mdl : Material.Model
    , actions : Maybe Actions
    , error : Maybe String
    }


init : Model
init =
    { mdl = Material.model
    , actions = Nothing
    , error = Nothing
    }
