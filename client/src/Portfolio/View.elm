module Portfolio.View exposing (render)

import Common.Decimal exposing (renderDecimal)
import Common.Error as Error
import Common.Styles exposing (padding, textLeft, textRight, toMdlCss)
import Html exposing (..)
import Html.Attributes exposing (href, src, style)
import Html.Events exposing (onClick)
import Identicon exposing (identicon)
import Main.Context exposing (Context)
import Main.Msg exposing (Msg(..))
import Main.Routing exposing (tokenPath)
import Material.Options as Options
import Material.Table as Table
import Material.Typography as Typo
import Model.Portfolio exposing (Portfolio, Position)
import Portfolio.Model exposing (Model)


render : Context -> Model -> Html Msg
render ctx model =
    let
        ( userName, ethereumAddress ) =
            case ctx.user of
                Just user ->
                    ( user.name, user.ethereumAddress )

                Nothing ->
                    ( "", "" )
    in
    div [ style [ ( "text-align", "center" ) ] ]
        [ header []
            [ text <| userName ++ " ("
            , a [ onClick Main.Msg.UserLogout ] [ text "logout" ]
            , text ")"
            ]
        , hr [] []
        , header [] [ text "Your tokens" ]
        , case model.error of
            Just _ ->
                Error.renderMaybeError model.error

            Nothing ->
                case model.portfolio of
                    Nothing ->
                        div [] [ text "Loading..." ]

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
                    [ text "You have no tokens yet" ]
                ]

        True ->
            div [] <| List.map (renderRow model) portfolio.positions


renderRow : Model -> Position -> Html Msg
renderRow model balance =
    a
        [ style
            [ ( "display", "flex" )
            , ( "border", "1px solid #ddd" )
            , ( "padding", "6px" )
            , ( "color", "black" )
            , ( "text-decoration", "none" )
            , ( "border-radius", "8px" )
            , ( "margin-bottom", "15px" )
            ]
        , href <| tokenPath balance.tokenId
        ]
        [ div
            [ style
                [ ( "padding", "15px" )
                , ( "text-align", "center" )
                , ( "width", "60px" )
                , ( "height", "60px" )
                , ( "border", "1px solid #ddd" )
                , ( "border-radius", "50%" )
                ]
            ]
            [ -- img
              --     [ style
              --         [ ( "width", "50px" )
              --         , ( "height", "50px" )
              --         , ( "border-radius", "8px" )
              --         , ( "margin-bottom", "5px" )
              --         ]
              --     , src token.logo
              --     ]
              --     []
              identicon "30px" balance.tokenName
            ]
        , div
            [ style
                [ ( "flex-grow", "1" )
                , ( "padding", "8px" )
                ]
            ]
            [ div
                [ style
                    [ ( "flex-grow", "1" )
                    , ( "font-weight", "bold" )
                    , ( "display", "block" )
                    ]
                ]
                [ header
                    [ style [ ( "margin", "0 0 5px 0" ) ] ]
                    [ text balance.tokenName
                    ]
                ]
            , div []
                [ text "You have "
                , renderDecimalWithPrecision balance.balance 2
                ]
            ]
        ]


renderDecimalWithPrecision : String -> Int -> Html a
renderDecimalWithPrecision value precision =
    case String.contains "." value of
        False ->
            span [] [ text value ]

        True ->
            let
                ( p, decimals, decimalZeros ) =
                    splitDecimal (round precision value) "" ""
            in
            span []
                [ text p
                , span [ style [ ( "font-size", "80%" ) ] ]
                    [ text decimals
                    , span [ style [ ( "opacity", "0.5" ) ] ] [ text decimalZeros ]
                    ]
                ]


{-| Splits decimal in three parts as follows:
10.340420000 becomes
part1: 10. --- main number
part2: 34042 --- decimal value
part3: 00000 --- Tail decimal zeros
-}
splitDecimal : String -> String -> String -> ( String, String, String )
splitDecimal value decimals decimalZeros =
    let
        ( decimalZeros, remaining_ ) =
            getDecimalZeros value ""

        ( decimals, remaining ) =
            getDecimals remaining_ ""
    in
    ( remaining, decimals, decimalZeros )


getDecimalZeros : String -> String -> ( String, String )
getDecimalZeros value decimalZeros =
    case String.right 1 value of
        "0" ->
            getDecimalZeros (String.dropRight 1 value) (decimalZeros ++ "0")

        remaining ->
            ( decimalZeros, value )


getDecimals : String -> String -> ( String, String )
getDecimals value decimals =
    case String.right 1 value of
        "." ->
            ( decimals, value )

        digit ->
            getDecimals (String.dropRight 1 value) (digit ++ decimals)


round : Int -> String -> String
round precision value =
    case String.split "." value of
        [ number, decimals ] ->
            number
                ++ "."
                ++ (if String.length decimals > precision then
                        String.left precision decimals

                    else
                        decimals
                   )

        _ ->
            value
