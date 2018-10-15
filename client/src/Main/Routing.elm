module Main.Routing exposing
    ( Route(..)
    , homepagePath
    , loginPath
    , matchers
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
    | CreateTokenRoute
    | UserLoginRoute


matchers : Parser (Route -> a) a
matchers =
    oneOf
        [ map HomepageRoute top
        , map TokensRoute (s "tokens")
        , map PortfolioRoute (s "portfolio")
        , map CreateTokenRoute (s "new")
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


portfolioPath : String
portfolioPath =
    "#portfolio"


loginPath : String
loginPath =
    "#login"


signupPath : String
signupPath =
    "#signup"
