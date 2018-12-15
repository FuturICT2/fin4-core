module CreateAction.View exposing (render)

import Common.Styles exposing (toMdlCss)
import CreateAction.Model exposing (Model)
import CreateAction.Msg exposing (Msg(..))
import CreateAction.StepName as StepName
import CreateAction.StepSuccess as StepSuccess
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (onClick)
import Main.Context exposing (Context)
import Main.Routing exposing (tokensPath)
import Material
import Material.Button as Button
import Material.Chip as Chip
import Material.Icon as Icon
import Material.Options as Options exposing (css)
import Material.Spinner as Spinner
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
        [ header
            [ style
                [ ( "text-align", "center" )
                , ( "margin", "30px 0px" )
                ]
            ]
            [ text "Create an action" ]
        , div []
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
    let
        buttonContent =
            case model.isCreatingAction of
                True ->
                    div [ style [ ( "margin-top", "10px" ) ] ]
                        [ Spinner.spinner [ Spinner.active True ] ]

                False ->
                    text "Create"
    in
    case model.step of
        0 ->
            div [ style [ ( "text-align", "center" ) ] ]
                [ Button.render Mdl
                    [ 23 ]
                    model.mdl
                    [ Button.ripple
                    , Button.raised
                    , Options.onClick PostAction
                    , toMdlCss buttonStyle
                    ]
                    [ buttonContent ]
                ]

        _ ->
            div [] []


mainStyle =
    style
        [ ( "font-family", "Roboto" )
        , ( "font-style", "condensed" )
        , ( "color", "black" )
        ]


buttonStyle : Attribute a
buttonStyle =
    style
        [ ( "height", "50px" )
        , ( "border-radius", "8px" )
        , ( "width", "100%" )
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
