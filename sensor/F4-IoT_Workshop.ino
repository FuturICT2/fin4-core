#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>

// Network Parameters
// TO DO //
char ssid[] = "Write the Name of your Network here"; //SSID of your Wi-Fi router
char pass[] = "Write your Network password here"; //Password of your Wi-Fi router
char accessToken[] = "Get the Access Token from the platform";

char pingAddress[] = "http://www.finfour.net/wapi/oracle-ping"; // sensor ping end point
char verificationAddress[] = "http://www.finfour.net/wapi/asset-block/sensor-verify"; // Claim verification end point
bool VerificationSent = false;

// Computation Variables
int values[100] = {}; // Stores latest 100 Humidity-Sensor Values
int arrayCounter = 0;
int val = 0;          // Humidity-Sensor Value
int average = 0;      // Average over last 100 Humidity-Sensor Values
int noise = 0;        // Noise-Sensor Value

// Pin Setup
int pushButton = D1;
int sensorPin1 = A0; // Humidity-Sensor

// Control Variables
int isPinging = 1; 
int allowChange = 0;
int buttonPressed = 0;
int pressedBefore = 0; // Old Button value
int loopCounter = 0;   // Counter for specific offsets
int timeOffset = 500;  // Modulo for resetting loopCounter
int delayTime = 2;    // Delay every loop cycle

void setup()
{
  // Start Serial Monitor (integrated)
  Serial.begin(9600);
  delay(10);

  // Set up Led and Push Button  
  pinMode(pushButton, INPUT_PULLUP); // Pullup-Resistor needed for avoiding floating Pin
  pinMode(LED_BUILTIN, OUTPUT);
  pinMode(sensorPin1, INPUT);


  // Set initial value for Led (1 == Led off, 0 == Led on)
  digitalWrite(LED_BUILTIN, 1);

  
  // Connect to Wi-Fi network
  Serial.println();
  Serial.println();
  Serial.print("Connecting to...");
  Serial.println(ssid);

  WiFi.begin(ssid, pass);

  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  
  Serial.println("");
  Serial.println("Wi-Fi connected successfully");
  

  // Set Array entries for Average Computation to initial Value received from Humidity-Sensor
  val = analogRead(sensorPin1);
  for (int i = 0; i < 100; i++) {
    values[i] = val;
  }
}

void loop () {

  // Read data (analog) from Humidity-Sensor and digitaly from Noise-Sensor
  val = analogRead(sensorPin1);

  // Check if PushButton is pressed (1 == not pressed, 0 == pressed)
  buttonPressed = digitalRead(pushButton);

  // Allow Change if Button has been released since the last push
  if (buttonPressed != pressedBefore) {
    allowChange = 1;
  } else {
    allowChange = 0;
  }

  // Change state if PushButton is pressed and we allow a change of state
  if (!buttonPressed && allowChange) {
    isPinging = !isPinging;
  }

  // Calculate Average over last 100 values (cyclic Array)
  if (loopCounter % 10 == 0) {
    values[arrayCounter] = val;
    arrayCounter = ++arrayCounter % 100;
    average = sum(values) / 100;
  }

  // Executing once per timeOffset * delayTime
  if (!loopCounter) {
    
    // Print if Device is not connected to Network
    if (WiFi.status() != WL_CONNECTED) {
      Serial.println("!!!!  NOT CONNECTED  !!!!");
    }

    // Ping the Server if pinging is enabled
    if (isPinging) {
      ping();
    }
    
    // Logging
    Serial.print("Sensor giving back = ");
    Serial.println(val);
    Serial.print("Average = ");
    Serial.println(average);
  }

  // Update last Button state
  pressedBefore = buttonPressed;

  // Send verification if measured Value from Humidity-Sensor is too far away from Average and if not yet verified (Difference relative to value of Average)
  if (val < (average - 100*(float(average) / 450)) && not VerificationSent) {

    // Send Verification
    verifyClaim();

    // Logging
    Serial.println("************* CLAB ************* Plant has been watered *** token has been issued ******");
    Serial.println(val);

    // Turn on Led
    digitalWrite(LED_BUILTIN, 0);
    
    VerificationSent = 1;

    // Pause program execution after Verification has been sent
    delay(20000);
  }

  // Loop counter to measure and display values at specific times
  loopCounter = ++loopCounter % timeOffset;
  delay(delayTime);
  
}

// Ping Server
void ping() {
  
  // Declare object of class HTTPClient
  HTTPClient http;

  // Specify Ping destination
  http.begin(pingAddress);

  // Specify Content type and authenticate as Oracle
  http.addHeader("Content-Type", "application/json");

  // Send Request and receive http Code
  int httpCode = http.POST(String("{\"accessToken\": \"") + accessToken + String("\"}"));

  // Logging
  if (httpCode == 200) {
    Serial.println("\n Ping! ------ \n");
  } else {
    Serial.println("\n server not available \n");
  }
  
  // Close connection
  http.end();
}

// Send Verification of Claim to Server 
void verifyClaim() {

  // Declare object of class HTTPClient
  HTTPClient http;

  // Specify request destination
  http.begin(verificationAddress);

  // Specify content-type header and "log in" as Oracle
  http.addHeader("Content-Type", "application/json");

  // Actually send the request
  int httpCode = http.POST(String("{\"isAccepted\": true, \"accessToken\": \"") + accessToken + String("\"}"));

  // Get the response payload
  String payload = http.getString();

  // Logging
  Serial.print("httpCode = ");
  Serial.println(httpCode);
  Serial.print("payload is = ");
  Serial.println(payload);

  // Close connection
  http.end();
}

// Compute Sum of values in Array
int sum(int values[]) {
  int s = 0;
  for (int i = 0; i < 100; i++) {
    s += values[i];
  }
  return s;
}
