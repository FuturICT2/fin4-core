module Actions.View exposing (render)

import Actions.Model exposing (Model)
import Common.Decimal exposing (renderDecimal)
import Common.Error as Error
import Common.Styles exposing (padding, textLeft, textRight, toMdlCss)
import Html exposing (..)
import Html.Attributes exposing (style)
import Html.Events exposing (onClick)
import Main.Context exposing (Context)
import Main.Msg exposing (Msg(..))
import Material.Options as Options
import Material.Table as Table
import Material.Typography as Typo
import Model.Actions exposing (Action, Actions)


render : Context -> Model -> Html Msg
render ctx model =
    div []
        [ div
            [ style
                [ ( "text-align", "center" )
                , ( "padding-top", "80px" )
                ]
            ]
            [ text "No actions yet" ]
        ]
