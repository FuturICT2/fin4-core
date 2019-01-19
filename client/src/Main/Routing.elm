module Main.Routing exposing
    ( Route(..)
    , actionsPath
    , fastSignupPath
    , forgotPassPath
    , homepagePath
    , loginPath
    , matchers
    , newActionPath
    , newTokenPath
    , parseLocation
    , portfolioPath
    , signupPath
    , termsPath
    , tokensPath
    )

import Navigation exposing (Location)
import UrlParser exposing ((</>), Parser, int, map, oneOf, parseHash, s, string, top)


type Route
    = NotFoundRoute
    | HomepageRoute
    | TokensRoute
    | TokenRoute Int
    | PortfolioRoute
    | ActionsRoute
    | CreateTokenRoute
    | CreateActionRoute
    | UserLoginRoute
    | UserForgotRoute
    | UserForgotResetPassRoute Int String
    | UserSignupRoute
    | UserFastSignupRoute


matchers : Parser (Route -> a) a
matchers =
    oneOf
        [ map HomepageRoute top
        , map TokensRoute (s "tokens")
        , map TokenRoute (s "token" </> int)
        , map PortfolioRoute (s "portfolio")
        , map ActionsRoute (s "actions")
        , map CreateTokenRoute (s "new")
        , map CreateActionRoute (s "new-action")
        , map UserLoginRoute (s "login")
        , map UserSignupRoute (s "signup")
        , map UserFastSignupRoute (s "fsignup")
        , map UserForgotRoute (s "forgot")
        , map UserForgotResetPassRoute (s "forgot-reset-password" </> int </> string)
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


forgotPassPath : String
forgotPassPath =
    "#forgot"


fastSignupPath : String
fastSignupPath =
    "#fsignup"


termsPath : String
termsPath =
    "#site-terms"


privacyPath : String
privacyPath =
    "#site-privacy"
