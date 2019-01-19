module CreateToken.StepSuccess exposing (render)

import Common.Error exposing (renderHttpError)
import CreateToken.Model exposing (Model)
import CreateToken.Msg exposing (Msg(..))
import Html exposing (..)
import Html.Attributes exposing (..)
import Main.Context exposing (Context)
import Material
import Material.Card as Card
import Material.Options as Options exposing (css)
import Material.Textfield as Textfield


render : Context -> Model -> Html Msg
render ctx model =
    div
        [ style [ ( "text-align", "center" ) ] ]
        [ div
            [ style
                [ ( "30px", "0" )
                , ( "padding", "0 30px" )
                , ( "margin", "30px 0" )
                ]
            ]
            [ text "Token was created"
            ]
        ]
