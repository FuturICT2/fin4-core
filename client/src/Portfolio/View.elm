module Portfolio.View exposing (render, renderData, renderRow)

import Common.Decimal exposing (renderDecimal)
import Common.Error as Error
import Common.Styles exposing (padding, textLeft, textRight, toMdlCss)
import Html exposing (..)
import Html.Attributes exposing (style)
import Main.Context exposing (Context)
import Main.Msg exposing (Msg(..))
import Material.Options as Options
import Material.Table as Table
import Material.Typography as Typo
import Model.Portfolio exposing (Portfolio, Position)
import Portfolio.Model exposing (Model)


render : Context -> Model -> Html Msg
render ctx model =
    let
        portfolioValue =
            case model.portfolio of
                Just portfolio ->
                    portfolio.valueInUSD

                Nothing ->
                    "0.00"

        address =
            case ctx.user of
                Just user ->
                    user.address

                Nothing ->
                    ""
    in
    div [ style [ ( "padding-top", "15px" ) ] ]
        [ Options.styled p [ Typo.headline ] [ text "Portfolio" ]
        , case model.error of
            Just _ ->
                Error.renderMaybeError model.error

            Nothing ->
                case model.portfolio of
                    Nothing ->
                        div [] [ text "..." ]

                    Just portfolio ->
                        renderData ctx model portfolio
        ]


renderData : Context -> Model -> Portfolio -> Html Msg
renderData ctx model portfolio =
    case List.length portfolio.positions > 0 of
        False ->
            div []
                [ Options.styled p
                    [ Typo.caption ]
                    [ text "You have no open position yet, visit Funding to deposit some assits into your acocunt" ]
                ]

        True ->
            div []
                [ Table.table []
                    [ Table.thead []
                        [ Table.tr []
                            [ Table.th [ toMdlCss textLeft ] [ text "Token" ]
                            , Table.th [ toMdlCss textRight ] [ text "Balance" ]
                            ]
                        ]
                    , Table.tbody [] <| List.map renderRow portfolio.positions
                    ]
                ]


renderRow : Position -> Html Msg
renderRow position =
    Table.tr []
        [ Table.td [ toMdlCss textLeft ] [ text position.name ]
        , Table.td [ toMdlCss textRight ] [ renderDecimal position.balance ]
        ]
