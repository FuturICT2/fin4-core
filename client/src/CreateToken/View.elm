module CreateToken.View exposing (render)

import Common.Spinner as Spinner
import CreateToken.Model exposing (Model)
import CreateToken.Msg exposing (Msg(..))
import CreateToken.StepName as StepName
import CreateToken.StepShares as StepShares
import CreateToken.StepSuccess as StepSuccess
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (onClick)
import Main.Context exposing (Context)
import Main.Routing exposing (tokensPath)
import Material
import Material.Chip as Chip
import Material.Icon as Icon
import Material.Options as Options
import Material.Textfield as Textfield


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
        [ modalStyle, mainStyle ]
        [ div [ titleStyle ]
            [ span
                [ style
                    [ ( "float", "left" )
                    , ( "visibility", visibility )
                    ]
                ]
                [ Icon.view "chevron_left"
                    [ Icon.size24
                    , Options.onClick StepBack
                    ]
                ]
            , text "Tokenization"
            , a
                [ style
                    [ ( "float", "right" )
                    , ( "color", "black" )
                    ]
                , href tokensPath
                ]
                [ Icon.view "close"
                    [ Icon.size24
                    ]
                ]
            ]
        , case model.isCreatingToken of
            True ->
                div [ style [ ( "marging-top", "50px" ) ] ]
                    [ Spinner.render "Creating your token, pleae wait a moment"
                    ]

            False ->
                div [ style [ ( "padding", "15px" ) ] ]
                    [ case model.step of
                        0 ->
                            StepName.render ctx model

                        1 ->
                            StepShares.render ctx model

                        2 ->
                            StepSuccess.render ctx model

                        _ ->
                            div [] [ text "" ]
                    ]
        , renderActionButton model.step
        ]


renderActionButton step =
    case step of
        1 ->
            div []
                [ div [ totalCost ] [ text "total cost: 0.002 ETH" ]
                , button
                    [ btnStyle
                    , onClick PostToken
                    ]
                    [ text "CONFIRM" ]
                ]

        2 ->
            button
                [ btnStyle
                ]
                [ text "Done" ]

        _ ->
            button
                [ btnStyle
                , onClick StepForward
                ]
                [ text "NEXT" ]


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
        , ( "position", "fixed" )
        , ( "bottom", "0" )
        , ( "font-size", "20px" )
        , ( "color", "black" )
        ]


totalCost : Attribute a
totalCost =
    style
        [ ( "height", "50px" )
        , ( "padding", "15px" )
        , ( "background-color", "black" )
        , ( "border", "none" )
        , ( "text-align", "center" )
        , ( "width", "100%" )
        , ( "position", "fixed" )
        , ( "bottom", "50px" )
        , ( "color", "white" )
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
