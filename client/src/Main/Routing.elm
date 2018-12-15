module Main.Routing exposing
    ( Route(..)
    , actionsPath
    , homepagePath
    , loginPath
    , matchers
    , newActionPath
    , newTokenPath
    , parseLocation
    , portfolioPath
    , signupPath
    , tokensPath
    )

import Navigation exposing (Location)
import UrlParser exposing ((</>), Parser, map, oneOf, parseHash, s, string, top)


type Route
    = NotFoundRoute
    | HomepageRoute
    | TokensRoute
    | PortfolioRoute
    | ActionsRoute
    | CreateTokenRoute
    | CreateActionRoute
    | UserLoginRoute


matchers : Parser (Route -> a) a
matchers =
    oneOf
        [ map HomepageRoute top
        , map TokensRoute (s "tokens")
        , map PortfolioRoute (s "portfolio")
        , map ActionsRoute (s "actions")
        , map CreateTokenRoute (s "new")
        , map CreateActionRoute (s "new-action")
        , map UserLoginRoute (s "login")
        ]


parseLocation : Location -> Route
parseLocation location =
    case parseHash matchers location of
        Just route ->
            route

        Nothing ->
            NotFoundRoute


homepagePath : String
homepagePath =
    "#"


tokensPath : String
tokensPath =
    "#tokens"


trendingPath : String
trendingPath =
    "#tokens"


newTokenPath : String
newTokenPath =
    "#new"


newActionPath : String
newActionPath =
    "#new-action"


portfolioPath : String
portfolioPath =
    "#portfolio"


actionsPath : String
actionsPath =
    "#actions"


loginPath : String
loginPath =
    "#login"


signupPath : String
signupPath =
    "#signup"
