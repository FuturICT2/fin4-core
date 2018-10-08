module TokenCreator.View exposing (render)

import Common.Decimal exposing (renderDecimal)
import Common.Error as Error
import Common.Styles exposing (padding, textLeft, textRight, toMdlCss)
import Html exposing (..)
import Html.Attributes exposing (style)
import Main.Context exposing (Context)
import Main.Msg exposing (Msg(..))
import Material.Button as Button
import Material.Elevation as Elevation
import Material.Options as Options
import Material.Table as Table
import Material.Textfield as Textfield
import Material.Typography as Typo
import TokenCreator.Model exposing (Model)
import TokenCreator.Msg


render : Context -> Model -> Html Msg
render ctx model =
    div [ style [ ( "padding-top", "15px" ) ] ]
        [ Options.styled p [ Typo.headline ] [ text "New Token" ]
        , div []
            [ div []
                [ Textfield.render Mdl
                    [ 1 ]
                    model.mdl
                    [ Textfield.label "Token name"
                    , Textfield.floatingLabel
                    , Textfield.value model.name
                    , Options.onInput TokenCreator.Msg.SetName
                    ]
                    []
                ]
            , div []
                [ Textfield.render Mdl
                    [ 2 ]
                    model.mdl
                    [ Textfield.label "Token total supply"
                    , Textfield.floatingLabel
                    , Textfield.value model.totalSupply
                    , Options.onInput TokenCreator.Msg.SetTotalSupply
                    ]
                    []
                ]
            , div []
                [ Button.render Mdl
                    [ 3 ]
                    model.mdl
                    [ Button.ripple
                    , Button.colored
                    , Button.raised
                    ]
                    [ text "Create" ]
                ]
            ]
        ]
