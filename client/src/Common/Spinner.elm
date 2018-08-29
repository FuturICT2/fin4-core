module Common.Spinner exposing (render)

import Html.Attributes exposing (..)
import Material.Spinner as Spinner
import Html exposing (..)


render msg =
    div [ style [ ( "padding", "10px" ) ] ]
        [ Spinner.spinner [ Spinner.active True ]
        , div
            [ style [ ( "padding", "5px" ) ]
            ]
            [ text msg ]
        ]
