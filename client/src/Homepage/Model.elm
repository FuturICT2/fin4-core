module Homepage.Model exposing (Model, init)

import Material
import Timeline.Model exposing (TimelineType(..))


type alias Model =
    { mdl : Material.Model
    , timeline : Timeline.Model.Model
    }


init : Model
init =
    { mdl = Material.model
    , timeline = Timeline.Model.init HomepageTimeline
    }
