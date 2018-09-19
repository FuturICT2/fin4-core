module Homepage.Homepage exposing (Msg(..), render)

import Html exposing (..)
import Html.Attributes exposing (..)
import Material
import Material.Elevation as Elevation
import Material.Options as Options exposing (css)
import Material.Typography as Typo


type Msg
    = Mdl (Material.Msg Msg)


render ctx =
    div [ style [ ( "margin-top", "15px" ) ] ]
        [ Options.styled p
            [ Typo.display2 ]
            [ text "Finance 4.0" ]
        ]
