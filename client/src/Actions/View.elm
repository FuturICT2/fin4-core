module Actions.View exposing (render)

import Actions.Model exposing (Model)
import Actions.Msg exposing (Msg(..))
import Common.Decimal exposing (renderDecimal, renderDecimalWithPrecision)
import Common.Error as Error
import Common.Styles exposing (padding, textLeft, textRight, toMdlCss)
import Html exposing (..)
import Html.Attributes exposing (style)
import Html.Events exposing (onClick)
import Identicon exposing (identicon)
import Main.Context exposing (Context)
import Material.Button as Button
import Material.Chip as Chip
import Material.Icon as Icon
import Material.Options as Options exposing (css)
import Material.Table as Table
import Material.Typography as Typo
import Model.Actions exposing (Action, Actions)


render : Context -> Model -> Html Msg
render ctx model =
    let
        content =
            case model.actions of
                Just actions ->
                    case List.length actions.entries of
                        0 ->
                            renderEmpty

                        _ ->
                            renderActions ctx model actions

                Nothing ->
                    renderEmpty
    in
    content


renderActions : Context -> Model -> Actions -> Html Msg
renderActions ctx model actions =
    div [ style [ ( "text-align", "center" ) ] ]
        [ header [] [ text "Actions" ]
        , div [] <| List.map (renderRow model) actions.entries
        ]


renderRow model action =
    let
        ( status, statusBG, statusColor ) =
            case action.status of
                0 ->
                    ( "new", "green", "white" )

                1 ->
                    ( "unconfirmed", "orange", "white" )

                2 ->
                    ( "done", "grey", "black" )

                _ ->
                    ( "unkown", "none", "black" )

        endsIn =
            case action.endsInHours of
                "1" ->
                    "ends in " ++ action.endsInMinutes ++ "mins"

                "0" ->
                    "ends in " ++ action.endsInMinutes ++ "mins"

                "-0" ->
                    ""

                _ ->
                    "ends in " ++ action.endsInHours ++ "hrs"
    in
    div
        [ style
            [ ( "border", "1px solid #ddd" )
            , ( "margin-bottom", "30px" )
            , ( "border-radius", "8px" )
            ]
        ]
        [ Chip.span
            [ css "background" statusBG
            , css "color" statusColor
            ]
            [ Chip.content [] [ text status ] ]
        , div
            [ style
                [ ( "padding", "5px 0 0 5px" )
                ]
            ]
            [ div
                [ style
                    [ ( "padding", "8px" )
                    ]
                ]
                [ header
                    [ style
                        [ ( "margin-bottom", "15px" )
                        , ( "line-height", "1.2" )
                        ]
                    ]
                    [ text action.description
                    ]
                ]
            ]
        , span [] [ text <| "you will get " ]
        , span [] <| List.map renderReward action.supporters
        , br [] []
        , span [] [ text <| "Fund this action" ]
        , br [] []
        , div
            [ style [ ( "margin", "15px 0" ) ] ]
            [ Button.render Mdl
                [ action.id, -1 ]
                model.mdl
                [ Button.ripple
                , Button.raised
                , Button.colored
                , css "margin-right" "15px"
                , Options.onClick (AddRewards action.id "1")
                ]
                [ text <| "+1 " ++ "GDPG"
                ]
            , Button.render Mdl
                [ action.id, -2 ]
                model.mdl
                [ Button.ripple
                , Button.raised
                , Button.colored
                , Options.onClick (AddRewards action.id "5")
                ]
                [ text <| "+5 " ++ "GDPG"
                ]
            ]
        , case action.isTimeLimit of
            True ->
                text endsIn

            False ->
                span [] []
        , div
            [ actionButtonsStyle
            ]
            [ Button.render Mdl
                [ action.id ]
                model.mdl
                [ Button.raised
                , Button.ripple
                , toMdlCss buttonStyle
                ]
                [ text "do it"

                -- , text <| " " ++ toString token.favouriteCount
                ]
            ]
        ]


renderReward reward =
    div []
        [ renderDecimalWithPrecision reward.amount 2
        , span [] [ text <| " " ++ reward.tokenName ++ " from " ]
        , span [] [ b [] [ text reward.userName ] ]
        ]


renderEmpty =
    div []
        [ div
            [ style
                [ ( "text-align", "center" )
                , ( "padding-top", "80px" )
                ]
            ]
            [ text "No actions yet" ]
        ]


actionButtonsStyle =
    style
        [ ( "border-top", "1px solid #ddd" )
        , ( "background", "#f2f2f2" )
        , ( "font-weight", "bold" )
        , ( "display", "block" )
        ]


buttonStyle =
    style
        [ ( "width", "100%" )
        , ( "text-align", "center" )
        , ( "line-height", "normal" )
        , ( "text-decoration", "none" )
        , ( "color", "inherit" )
        , ( "display", "inline-block" )
        , ( "box-shadow", "none" )
        , ( "font-size", "26px" )
        , ( "padding", "15px" )
        , ( "height", "60px" )
        ]
