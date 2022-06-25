# How Google Uses Go

[Using Go at Google](https://go.dev/solutions/google/)

- In 2015, to keep up with Google’s scale, Google’s Core Data Solutions team needed to rewrite the indexing stack from a single monolithic binary written in C++ to multiple components in a microservices architecture. The team decided to rewrite many indexing services in Go, which they now use to power the majority of their architecture. As a result, Google’s web indexing was re-architected within a year. More impressively, most developers on the team were rewriting in Go while also learning it.
- Behind the scenes, Chrome has an extensive fleet of backends. Among these is the Chrome Optimization Guide service. This service forms an important basis for Chrome’s user experience strategy, operating in the critical path for users, and is implemented in Go. Millions of users rely on this service to make their Chrome experience better. When the Chrome engineering team started building the service, only a few members had comfort with Go.
- Go is the language of choice for Google SRE teams. The majority of Google production is managed and maintained by our systems written in Go. Go’s simplicity means that the code is easy to follow, whether it is to spot bugs during review or when trying to determine exactly what happened during a service disruption.
- Go is used by the Google Cloud Platform team for solutions such as Anthos, GKE, and Istio
- The Firebase Hosting team has replaced 100% of backend Node.js code with Go. Hundreds of thousands of customers host their websites with Firebase Hosting, which means Go code is used to serve billions of requests per day. 
- Go is generally very readable and understandable. The language’s error handling, receivers, and interfaces are all easy to understand due to the idioms in the language.


