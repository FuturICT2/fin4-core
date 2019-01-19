module Main.NotFound exposing (render)

import Html exposing (..)
import Main.Model exposing (Model)
import Main.Msg exposing (Msg)


render : Model -> Html Msg
render model =
    div [] [ text "not found 404" ]
