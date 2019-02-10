module Homepage.View exposing (render)

import Homepage.Model exposing (Model)
import Homepage.Msg exposing (Msg(..))
import Homepage.NotAuth
import Html exposing (..)
import Html.Attributes exposing (style)
import Main.Context exposing (Context)
import Timeline.View


render : Context -> Model -> Html Msg
render ctx model =
    case ctx.user of
        Just data ->
            Html.map TimelineMsg <|
                Timeline.View.render ctx model.timeline

        Nothing ->
            Homepage.NotAuth.render ctx model
