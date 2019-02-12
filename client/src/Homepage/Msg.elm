module Homepage.Msg exposing (Msg(..))

import Material
import Timeline.Msg


type Msg
    = Mdl (Material.Msg Msg)
    | TimelineMsg Timeline.Msg.Msg
