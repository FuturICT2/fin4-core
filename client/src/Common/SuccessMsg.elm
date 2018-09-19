module Common.SuccessMsg exposing (render)

import Html.Attributes exposing (..)
import Html exposing (..)
import Material.Options as Options
import Material.Typography as Typo


render msg =
    div
        [ style
            [ ( "padding", "20px 0" )
            , ( "color", "rgb(33, 150, 243)" )
            ]
        ]
        [ Options.styled p [ Typo.body2 ] [ text msg ] ]
