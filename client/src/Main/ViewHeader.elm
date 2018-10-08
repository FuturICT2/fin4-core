module Main.ViewHeader exposing
    ( authHeader
    , authRightSideNav
    , bhNavBorderedRightStyle
    , bhNavItemStyle
    , bhStyle
    , brandColor
    , notAuthHeader
    , notAuthRightSideNav
    , render
    , renderForDesktop
    , renderForMobile
    , textColor
    , thLogoStyle
    , thNavItemStyle
    , thNavStyle
    , thStyle
    , thSubNavItemStyle
    )

import Html exposing (..)
import Html.Attributes exposing (attribute, href, style)
import Html.Events exposing (onClick)
import Main.Model exposing (Model)
import Main.Msg exposing (Msg(..))
import Main.Routing
    exposing
        ( Route(..)
        , homepagePath
        , newTokenPath
        , portfolioPath
        , tokensPath
        )
import Material.Icon as Icon


render : Model -> Html Msg
render model =
    if model.context.window.isMobile then
        renderForMobile model

    else
        renderForDesktop model


renderForDesktop : Model -> Html Msg
renderForDesktop model =
    div
        [ thStyle ]
        (case model.context.user of
            Just _ ->
                authHeader model

            _ ->
                notAuthHeader model
        )


authHeader : Model -> List (Html Msg)
authHeader model =
    [ authRightSideNav model
    , a [ thNavItemStyle, href tokensPath ] [ text "Tokens" ]
    , a [ thNavItemStyle, href portfolioPath ] [ text "Balances" ]
    , a [ thNavItemStyle, href <| homepagePath ] [ text "New Token" ]
    ]


notAuthHeader : Model -> List (Html Msg)
notAuthHeader model =
    [ notAuthRightSideNav model
    , a [ href homepagePath, thNavItemStyle ] [ text "Fin4" ]
    , a
        [ href tokensPath
        , thNavItemStyle
        ]
        [ text "Tokens" ]
    , a
        [ href portfolioPath
        , thNavItemStyle
        ]
        [ text "Balances" ]
    , a
        [ href tokensPath
        , thNavItemStyle
        ]
        [ text "Actions" ]
    , a
        [ href tokensPath
        , thNavItemStyle
        ]
        [ text "New Token" ]
    ]


renderForMobile : Model -> Html Msg
renderForMobile model =
    case model.context.user of
        Just _ ->
            div [ bhStyle ]
                [ div []
                    [ a
                        [ href tokensPath
                        , bhNavItemStyle
                        , bhNavBorderedRightStyle
                        , style [ ( "width", "25%" ) ]
                        ]
                        [ Icon.view "timeline"
                            [ Icon.size24
                            ]
                        ]
                    , a
                        [ href newTokenPath
                        , bhNavItemStyle
                        , bhNavBorderedRightStyle
                        , style [ ( "width", "25%" ) ]
                        ]
                        [ Icon.view "add_circle_outline"
                            [ Icon.size24
                            ]
                        ]
                    , a
                        [ href portfolioPath
                        , bhNavItemStyle
                        , bhNavBorderedRightStyle
                        , style [ ( "width", "25%" ) ]
                        ]
                        [ Icon.view "person"
                            [ Icon.size24
                            ]
                        ]
                    , a
                        [ href portfolioPath
                        , bhNavItemStyle
                        , bhNavBorderedRightStyle
                        , style [ ( "width", "25%" ) ]
                        ]
                        [ Icon.view "settings"
                            [ Icon.size24
                            ]
                        ]
                    ]
                ]

        _ ->
            div [ bhStyle ]
                [ a
                    [ href tokensPath
                    , bhNavItemStyle
                    , bhNavBorderedRightStyle
                    , style [ ( "width", "25%" ) ]
                    ]
                    [ text "Tokens" ]
                , a
                    [ href portfolioPath
                    , bhNavItemStyle
                    , bhNavBorderedRightStyle
                    , style [ ( "width", "25%" ) ]
                    ]
                    [ text "Balances" ]
                , a
                    [ href homepagePath
                    , bhNavItemStyle
                    , bhNavBorderedRightStyle
                    , style [ ( "width", "25%" ) ]
                    ]
                    [ text "Actions" ]
                , a
                    [ href homepagePath
                    , bhNavItemStyle
                    , bhNavBorderedRightStyle
                    , style [ ( "width", "25%" ) ]
                    ]
                    [ text "New Token" ]
                ]


notAuthRightSideNav : Model -> Html Msg
notAuthRightSideNav model =
    div [ thNavStyle ]
        [ a [ href "#", thSubNavItemStyle ] [ text "Activate MetaMask" ] ]


authRightSideNav : Model -> Html Msg
authRightSideNav model =
    div [ thNavStyle ]
        [ a
            [ thSubNavItemStyle
            ]
            [ text "Logout" ]
        ]



-- Common


brandColor : String
brandColor =
    "#3f51b5"



--"#262b3e"
--"#131722"
--"rgb(33,150,243)"


textColor : String
textColor =
    "#e1ecf2"



--"#fff"
-- Desktop styles: TopHeader : th


thStyle : Attribute a
thStyle =
    style
        [ ( "height", "45px" )
        , ( "padding-top", "10px" )
        , ( "box-sizing", "border-box" )
        , ( "padding-left", "10px" )
        , ( "padding-right", "10px" )
        , ( "line-height", "25px" )
        , ( "background", brandColor )
        , ( "width", "100%" )
        , ( "position", "fixed" )
        , ( "top", "0" )
        , ( "z-index", "1000" )
        ]


thNavStyle : Attribute a
thNavStyle =
    style
        [ ( "float", "right" )
        ]


thNavItemStyle : Attribute a
thNavItemStyle =
    style
        [ ( "margin", "0 8px" )
        , ( "text-decoration", "none" )
        , ( "font-size", "14px" )
        , ( "letter-spacing", ".7px" )
        , ( "color", textColor )
        ]


thSubNavItemStyle : Attribute a
thSubNavItemStyle =
    style
        [ ( "padding", "10px" )
        , ( "text-decoration", "none" )
        , ( "font-size", "12px" )
        , ( "font-weight", "normal" )
        , ( "color", textColor )
        ]


thLogoStyle : Attribute a
thLogoStyle =
    style
        [ ( "text-transform", "uppercase" )
        , ( "text-decoration", "none" )
        , ( "font-size", "12px" )
        , ( "margin", "0px 40px 0px 25px" )
        , ( "color", textColor )
        ]



-- Mobile Styles, BottomHeader :bh


bhStyle : Attribute a
bhStyle =
    style
        [ ( "position", "fixed" )
        , ( "bottom", "0" )
        , ( "left", "0" )
        , ( "width", "100%" )
        , ( "height", "40px" )
        , ( "background", "white" )
        , ( "min-width", "320px" )
        , ( "opacity", "1" )
        , ( "z-index", "100" )
        ]


bhNavItemStyle : Attribute a
bhNavItemStyle =
    style
        [ ( "display", "inline-block" )
        , ( "text-align", "center" )
        , ( "line-height", "50px" )
        , ( "text-decoration", "none" )
        , ( "color", "black" )
        , ( "overflow", "hidden" )
        , ( "cursor", "pointer" )
        ]


bhNavBorderedRightStyle : Attribute a
bhNavBorderedRightStyle =
    style
        [ ( "box-sizing", "border-box" )
        , ( "border-right", "1px solid " ++ textColor )
        ]
