module CreateAsset.View exposing (render)

import Common.Error exposing (renderHttpError)
import Common.Styles exposing (marginTop, paddingBottom, paddingTop, textCenter, toMdlCss)
import CreateAsset.Model exposing (Model)
import CreateAsset.Msg exposing (Msg(..))
import CreateAsset.Types exposing (..)
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (onClick, onInput)
import Identicon exposing (identicon)
import Main.Context exposing (Context)
import Main.Routing exposing (assetPath, exploreAssetsPath)
import Material.Button as Button
import Material.Options as Options
import Material.Spinner
import Material.Toggles as Toggles
import Material.Typography as Typography


render : Context -> Model -> Html Msg
render ctx model =
    div []
        [ case model.step of
            InformationView ->
                renderInformation ctx model

            FormView ->
                renderForm ctx model

            SuccessView ->
                renderSuccess ctx model
        ]


renderInformation : Context -> Model -> Html Msg
renderInformation ctx model =
    div [ informationStyle ]
        [ div [ informationHelpStyle ]
            [ Options.styled p
                [ Typography.body1 ]
                [ text """Once you create a topic, a token will
                      be created for that topic with 0 currency supply. The feedback from people on
                      this topic initiates and generates the cryptocurrency
                      supply. Therefore, the more content and posts are added to
                      the topic, the more cryptocurrency is mined "socially"."""
                ]
            ]
        , Button.render Mdl
            [ 1 ]
            model.mdl
            [ Button.raised
            , Button.ripple
            , Button.colored
            , Options.onClick (SetActiveView FormView)
            ]
            [ text "next"
            ]
        ]


renderForm : Context -> Model -> Html Msg
renderForm ctx model =
    div
        [ mainStyle ]
        [ div [ identiconStyle ] [ identicon "60px" model.name ]
        , div [ marginTop 30 ]
            [ p [ hintStle ] [ text "Topic name" ]
            , input
                [ inputStyle
                , placeholder "e.g Trees"
                , value model.name
                , onInput SetName
                ]
                []
            ]
        , div []
            [ p [ hintStle ] [ text "Topic currency symbol" ]
            , input
                [ inputStyle
                , placeholder "e.g TRR"
                , value model.symbol
                , onInput SetSymbol
                ]
                []
            ]
        , div
            []
            [ p [ hintStle ] [ text "Topic description" ]
            , textarea
                [ textareaStyle
                , value model.purpose
                , placeholder "e.g Plant a tree and get TRR coins!"
                , onInput SetDescription
                , rows 5
                ]
                []
            ]
        , Toggles.checkbox Mdl
            [ 0 ]
            model.mdl
            [ Options.onToggle ToggleOracleType
            , Toggles.ripple
            , Toggles.value model.isSensor
            ]
            [ text "Check this if the oracle is an IoT sensor" ]
        , renderHttpError model.createAssetError
        , div [ style [ ( "text-align", "center" ) ] ]
            [ Button.render Mdl
                [ 0 ]
                model.mdl
                [ Button.raised
                , Button.colored
                , Options.onClick PostAsset
                , Options.disabled model.isCreatingAsset
                , toMdlCss (marginTop 15)
                ]
                [ case model.isCreatingAsset of
                    True ->
                        Material.Spinner.spinner
                            [ Material.Spinner.active True ]

                    False ->
                        text "Confirm"
                ]
            ]
        ]


renderSuccess : Context -> Model -> Html Msg
renderSuccess ctx model =
    div [ paddingTop 45, textCenter ]
        [ header [ paddingBottom 45 ] [ text "Congratulations" ]
        , div [ paddingBottom 45 ]
            [ text "Your topic has been successfully created"
            ]
        , a [ href (assetPath model.createdAssetId) ]
            [ text <|
                "Go to "
                    ++ model.createdAssetName
                    ++ " ("
                    ++ model.createdAssetSymbol
                    ++ ")"
            ]
        , div [ paddingBottom 30, paddingTop 30 ] [ text "OR" ]
        , a [ href exploreAssetsPath ] [ text "Explore other tokens" ]
        ]


mainStyle : Attribute a
mainStyle =
    style [ ( "margin", "15px 15px 60px 15px" ) ]


inputStyle : Attribute a
inputStyle =
    style
        [ ( "border-radius", "8px" )
        , ( "background-color", "#f2f2f2" )
        , ( "height", "50px" )
        , ( "border", "none" )
        , ( "padding", "15px" )
        , ( "width", "100%" )
        , ( "box-sizing", "border-box" )
        , ( "margin-bottom", "15px" )
        ]


textareaStyle : Attribute a
textareaStyle =
    style
        [ ( "border-radius", "8px" )
        , ( "background-color", "#f2f2f2" )
        , ( "border", "none" )
        , ( "padding", "15px" )
        , ( "width", "100%" )
        , ( "box-sizing", "border-box" )
        , ( "margin-bottom", "15px" )
        ]


identiconStyle : Attribute a
identiconStyle =
    style
        [ ( "width", "100px" )
        , ( "height", "100px" )
        , ( "border", "1px solid #ddd" )
        , ( "border-radius", "50%" )
        , ( "margin", "auto" )
        , ( "padding", "20px" )
        ]


informationStyle : Attribute a
informationStyle =
    style
        [ ( "text-align", "center" )
        , ( "padding-top", "45px" )
        ]


informationHelpStyle : Attribute a
informationHelpStyle =
    style
        [ ( "padding", "45px 10px" )
        ]


hintStle : Attribute a
hintStle =
    style
        [ ( "margin", " 0" )
        , ( "padding-left", "2px" )
        ]
