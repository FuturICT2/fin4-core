module Person.View exposing (render)

import Common.Decimal exposing (renderDecimalWithPrecision)
import Common.Error exposing (renderHttpError)
import Common.Spinner as Spinner
import Common.Styles exposing (toMdlCss)
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (on, onClick, onInput)
import Identicon exposing (identicon)
import Json.Decode as JD
import Main.Context exposing (Context)
import Main.Routing exposing (tokenPath)
import Material
import Material.Button as Button
import Material.Icon as Icon
import Material.Options as Options exposing (css)
import Material.Table as Table
import Material.Tabs as Tabs
import Model.Persons exposing (MinedToken, Person)
import Person.Model exposing (Model)
import Person.Msg exposing (Msg(..))


render : Context -> Model -> Int -> Html Msg
render ctx model personId =
    div [ style [ ( "text-align", "center" ), ( "margin-bottom", "60px" ) ] ]
        [ case model.data of
            Nothing ->
                renderSpinner

            Just person ->
                renderPerson ctx model person
        ]


renderSpinner : Html Msg
renderSpinner =
    div []
        [ Spinner.render "One moment..." ]


renderPerson : Context -> Model -> Person -> Html Msg
renderPerson ctx model person =
    div []
        [ header [] [ text person.name ]
        , div [] [ text "User tokens" ]
        , renderTokens ctx model person
        ]



-- renderTabs : Context -> Model -> Person -> Html Msg
-- renderTabs ctx model person =
--     Tabs.render Mdl
--         [ 0 ]
--         model.mdl
--         [ Tabs.ripple
--         , Tabs.onSelectTab SelectTab
--         , Tabs.activeTab model.tab
--         ]
--         [ Tabs.label
--             [ Options.center ]
--             [ text <| "Claims "
--             ]
--         , Tabs.label
--             [ Options.center ]
--             [ text <| "Tokens "
--             ]
--         ]
--         [ case model.tab of
--             0 ->
--                 div [] [ text "claims" ]
--
--             _ ->
--                 renderTokens ctx model person
--         ]


renderTokens : Context -> Model -> Person -> Html Msg
renderTokens ctx model person =
    case List.length person.minedTokens of
        0 ->
            div
                [ style
                    [ ( "padding", "15px" )
                    ]
                ]
                [ text "There are no tokens yet"
                ]

        _ ->
            Table.table
                []
                [ Table.thead []
                    [ Table.tr []
                        [ Table.th [] [ text "Token" ]
                        , Table.th [] [ text "Balance" ]
                        , Table.th [] [ text "Mined" ]
                        ]
                    ]
                , Table.tbody [] <|
                    List.map (renderToken model) person.minedTokens
                ]


renderToken : Model -> MinedToken -> Html Msg
renderToken model token =
    Table.tr []
        [ Table.td
            []
            [ a [ href <| tokenPath token.tokenId ] [ text token.tokenName ] ]
        , Table.td
            []
            [ renderDecimalWithPrecision token.balance 2
            , small [] [ text token.tokenSymbol ]
            ]
        , Table.td
            []
            [ renderDecimalWithPrecision token.mined 2
            , small [] [ text token.tokenSymbol ]
            ]
        ]
