module Homepage.Command exposing (commands)

import Homepage.Model exposing (Model)
import Homepage.Msg exposing (Msg(..))
import Main.Context exposing (Context)
import Timeline.Command
import Timeline.Model exposing (TimelineType(..))


commands : Context -> Model -> Cmd Msg
commands ctx model =
    Cmd.map TimelineMsg <| Timeline.Command.commands ctx HomepageTimeline
