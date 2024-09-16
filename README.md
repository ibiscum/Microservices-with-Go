# Microservices with Go

[![build](https://github.com/ibiscum/Microservices-with-Go/actions/workflows/build.yml/badge.svg)](https://github.com/ibiscum/Microservices-with-Go/actions/workflows/build.yml)
[![lint](https://github.com/ibiscum/Microservices-with-Go/actions/workflows/lint.yml/badge.svg)](https://github.com/ibiscum/Microservices-with-Go/actions/workflows/lint.yml)

<a href="https://www.packtpub.com/product/microservices-with-go/9781804617007"><img src="https://m.media-amazon.com/images/I/412x+RC-FJL._SX403_BO1,204,203,200_.jpg" alt="Microservices with Go" height="256px" align="right"></a>

This is the code repository for [Microservices with Go](https://www.packtpub.com/product/microservices-with-go/9781804617007), published by Packt.

**Building scalable and reliable microservices with Go**

## What is this book about?

This book covers the following exciting features:
* Get familiar with the industry’s best practices and solutions in microservice development
* Understand service discovery in the microservices environment
* Explore reliability and observability principles
* Discover best practices for asynchronous communication
* Focus on how to write high-quality unit and integration tests in Go applications
* Understand how to profile Go microservices

If you feel this book is for you, get your [copy](https://www.amazon.com/dp/1804617008) today!

<a href="https://www.packtpub.com/?utm_source=github&utm_medium=banner&utm_campaign=GitHubBanner"><img src="https://raw.githubusercontent.com/PacktPublishing/GitHub/master/GitHub.png"
alt="https://www.packtpub.com/" border="5" /></a>

## Errata

* Page 39: In the Handler section, the import statement "github.com/ibiscum/Microservices-with-Go/Chapter0X/rating/internal/controller" must be read as "github.com/ibiscum/Microservices-with-Go/Chapter0X/rating/internal/controller/rating".
* Page 64-67: The types "serviceName" and "instanceID" are defined, but these same names are also used as method parameters, causing a naming clash as the methods Register, Deregister, ReportHealthyState and ServiceAddresses also access the serviceAddrs slice which uses serviceName and instanceID as types also.
* Page 69: The for loop on this page must be read as " for _, e := range entries {        res = append(res, fmt.Sprintf("%s:%d", e.Service.Address, e.Service.Port)))    }". That is "res = append(res, " is repeated twice which must be ignored.


## Suggestion -- Chapter07: quick fixes for MySQL, grpcurl

In case you are not able to follow along due to some steps being left out, please try these steps before chapter07 examples:
* run ```CREATE DATABASE movieexample``` (you probably want to access the container instance to execute this command by using ```mysql -uroot -ppassword``` in Docker first)
* run ```go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest``` to get grpcurl on your machine

Suggested by our reader <b>jasonsalas</b> [here](https://github.com/PacktPublishing/Microservices-with-Go/issues/9).

## Instructions and Navigations
All of the code is organized into folders. For example, Chapter02.

The code will look like the following:
```
package main
import (
    “encoding/json”
    “fmt”
    “os”
    “time”

    “github.com/confluentinc/confluent-kafka-go/v2/kafka”
    “github.com/ibiscum/Microservices-with-Go/Chapter0X/rating/pkg/model”
)
```

**Following is what you need for this book:**
This book is for all types of developers, from people interested in learning how to write microservices in Go to seasoned professionals who want to take the next step in mastering the art of writing scalable and reliable microservice-based systems. A basic understanding of Go will come in handy.

With the following software and hardware list you can run all code files present in the book (Chapter 1-13).
### Software and Hardware List
| Chapter | Software required | OS required |
| -------- | ------------------------------------ | ----------------------------------- |
| 1-13 | Go 1.11 or above | Windows, Mac OS X, and Linux (Any) |
| 1-13 | Docker | Windows, Mac OS X, and Linux (Any) |
| 1-13 | grpcurl | Windows, Mac OS X, and Linux (Any) |
| 1-13 | Kubernetes | Windows, Mac OS X, and Linux (Any) |
| 1-13 | Prometheus | Windows, Mac OS X, and Linux (Any) |
| 1-13 | Jaeger | Windows, Mac OS X, and Linux (Any) |
| 1-13 | Graphviz | Windows, Mac OS X, and Linux (Any) |

We also provide a PDF file that has color images of the screenshots/diagrams used in this book. [Click here to download it](https://packt.link/1fb2C).

### Related products
* Go for DevOps [[Packt]](https://www.packtpub.com/product/go-for-devops/9781801818896?utm_source=github&utm_medium=repository&utm_campaign=9781801818896) [[Amazon]](https://www.amazon.com/dp/1801818894)

* Event-Driven Architecture in Golang [[Packt]](https://www.packtpub.com/product/event-driven-architecture-in-golang/9781803238012#:~:text=Event%2DDriven%20Architecture%20in%20Golang%20is%20an%20approach%20used%20to,internally%2C%20and%20externally%20using%20messages.?utm_source=github&utm_medium=repository&utm_campaign=9781803238012) [[Amazon]](https://www.amazon.com/dp/1803238011)


## Get to Know the Author
**Alexander Shuiskov**
Senior Software Engineer II at Uber. He carries immense experience in Distributed Systems, Microservices, and Observability. He has completed his MSc in Computer Science and after which he has worked in prominent companies like eBay, Booking, and so on. He has led the development of Uber alerting and contributed to 25 Uber services. He was also an active contributor to the development of the eBay Ad platform.

### Download a free PDF
<i>If you have already purchased a print or Kindle version of this book, you can get a DRM-free PDF version at no cost.<br>Simply click on the link to claim your free PDF.</i>
<p align="center"> <a href="https://packt.link/free-ebook/9781804617007">https://packt.link/free-ebook/9781804617007 </a> </p>
