module CreateToken.View exposing (render)

import CreateToken.Model exposing (Model)
import CreateToken.Msg exposing (Msg(..))
import CreateToken.StepName as StepName
import CreateToken.StepSuccess as StepSuccess
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (onClick)
import Main.Context exposing (Context)
import Main.Routing exposing (tokensPath)
import Material
import Material.Button as Button
import Material.Chip as Chip
import Material.Icon as Icon
import Material.Options as Options
import Material.Textfield as Textfield
import Material.Typography as Typo


render : Context -> Model -> Html Msg
render ctx model =
    let
        visibility =
            case model.step of
                0 ->
                    "hidden"

                _ ->
                    "visible"
    in
    div
        [ mainStyle ]
        [ h3 [ style [ ( "text-align", "center" ) ] ] [ text "Tokenizer" ]
        , div [ style [ ( "margin", "15px" ) ] ]
            [ case model.step of
                0 ->
                    StepName.render ctx model

                1 ->
                    StepSuccess.render ctx model

                _ ->
                    div [] [ text "" ]
            ]
        , renderActionButton model
        ]


renderActionButton model =
    case model.step of
        0 ->
            div [ style [ ( "text-align", "center" ) ] ]
                [ div [ totalCost ] [ text "network fee: 0.002 ETH" ]
                , Button.render Mdl
                    [ 23 ]
                    model.mdl
                    [ Button.ripple
                    , Button.raised
                    , Options.onClick PostToken
                    ]
                    [ text "Create Token" ]
                ]

        _ ->
            div [] []


mainStyle =
    style
        [ ( "font-family", "Roboto" )
        , ( "font-style", "condensed" )
        , ( "color", "black" )
        ]


btnStyle : Attribute a
btnStyle =
    style
        [ ( "height", "50px" )
        , ( "background-color", "#f2f2f2" )
        , ( "background-color", "#dedede" )
        , ( "border", "none" )
        , ( "border-top", "1px solid black" )
        , ( "font-weight", "bold" )
        , ( "text-align", "center" )
        , ( "color", "black" )
        , ( "width", "100%" )
        , ( "font-size", "20px" )
        , ( "color", "black" )
        ]


totalCost : Attribute a
totalCost =
    style
        [ ( "height", "50px" )
        , ( "padding", "15px" )
        , ( "border", "none" )
        , ( "text-align", "center" )
        , ( "width", "100%" )
        ]


titleStyle =
    style
        [ ( "text-align", "center" )
        , ( "font-weight", "bold" )
        , ( "padding", "15px" )
        , ( "font-size", "20px" )
        ]


modalStyle : Attribute a
modalStyle =
    style
        [ ( "background-color", "rgba(255,255,255,1.0)" )
        , ( "position", "fixed" )
        , ( "top", "0" )
        , ( "bottom", "0" )
        , ( "left", "0" )
        , ( "right", "0" )
        , ( "padding", "0px" )
        , ( "z-index", "100000000" )
        ]
