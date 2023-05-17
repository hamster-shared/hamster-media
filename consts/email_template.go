package consts

const EmailTemplate = `<!DOCTYPE html>
<html>

<body>
  <p><strong>Application Information is as follows</strong></p>
  <p><strong>Application Time: </strong>{{.ApplicationTime}}</p>
  <p><strong>Name: </strong>{{.Name}}</p>
  <p><strong>Email: </strong>{{.Email}}</p>
  <p><strong>Social Account: </strong>{{.SocialAccount}}</p>
  <p><strong>Middleware Category: </strong>{{.MiddlewareCategory}}</p>
  <p><strong>Middleware Information: </strong>{{.MiddlewareInformation}}</p>
</body>

</html>`