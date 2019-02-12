module Homepage.Update exposing (update)

import Homepage.Model exposing (Model)
import Homepage.Msg exposing (Msg(..))
import Main.Context exposing (Context)
import Material
import Timeline.Update


update : Context -> Msg -> Model -> ( Model, Cmd Msg )
update ctx msg model =
    case msg of
        TimelineMsg msg_ ->
            let
                ( childModel, cmd ) =
                    Timeline.Update.update ctx msg_ model.timeline
            in
            { model | timeline = childModel } ! [ Cmd.map TimelineMsg cmd ]

        Mdl msg_ ->
            Material.update Mdl msg_ model
