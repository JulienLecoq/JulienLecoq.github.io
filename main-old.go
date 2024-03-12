package main

// func main() {
//     app.Route("/", &root.Root{})
//     app.RunWhenOnBrowser()
//
//     // Wrap the app.Handler to check for WASM file requests
//     wasmHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         // Check if the request is for a WASM file
//         if strings.HasSuffix(r.URL.Path, ".wasm") {
//             // Set the correct Content-Type for WASM files
//             w.Header().Set("Content-Type", "application/wasm")
//         }
//         // Use the existing app.Handler to serve the request
//         // This ensures that non-WASM files are served normally
//         &app.Handler{
//             Name:        "Hello",
//             Description: "An Hello World! example",
//             Scripts: []string{
//                 "/web/main.js",
//             },
//             Styles: []string{
//                 "/web/main.css",
//             },
//         }.ServeHTTP(w, r)
//     })
//
//     // Use the custom handler for serving requests
//     http.Handle("/", wasmHandler)
//
//     // Start the server
//     if err := http.ListenAndServe(":8000", nil); err != nil {
//         log.Fatal(err)
//     }
// }
