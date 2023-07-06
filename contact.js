const submitData = (e) => {
  let name = document.getElementById("input-name").value
  let email = document.getElementById("input-email").value
  let phone = document.getElementById("input-phone").value
  let subject = document.getElementById("input-subject").value
  let message = document.getElementById("input-message").value

  let object = {
    name,
    email,
    phone,
    subject,
    message,
  }

  console.log(object)

  if (name === "") {
    return alert("Please Fill The Name")
  } else if (email === "") {
    return alert("Please Fill The Email")
  } else if (phone === "") {
    alert("Please Fill The Phone Number")
  } else if (subject === "") {
    alert("Please Fill The Subject")
  } else if (message === "") {
    alert("Please Fill The Message")
  }

  const emailReceive = "inifryeyay@gmail.com"

  let a = document.createElement("a")
  a.href = `mailto:${emailReceive}?subject=${subject}&body=Halo nama saya ${name}, \n${message}, silahkan kontak : ${phone}`
  a.click()
}
