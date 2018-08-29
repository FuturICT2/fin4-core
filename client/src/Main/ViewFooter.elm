module Main.ViewFooter exposing (render)

import Html exposing (..)
import Html.Attributes exposing (href)
import Main.Model exposing (Model)
import Main.Msg exposing (Msg)
import Material.Footer as Footer
import Material.Options as Options


render : Model -> Html Msg
render model =
    Footer.mini [ Options.css "margin-top" "20px" ]
        { left =
            Footer.left []
                [ Footer.links []
                    []
                ]
        , right = Nothing
        }
