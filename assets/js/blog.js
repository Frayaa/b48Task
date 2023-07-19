let dataBlog = []

const addBlog = (event) => {
  event.preventDefault()

  let project = document.getElementById("input-project").value
  let startDate = document.getElementById("input-start").value
  let endDate = document.getElementById("input-end").value
  let input = document.querySelectorAll(".multi-input:checked")
  let description = document.getElementById("input-description").value
  let image = document.getElementById("input-image").files
  let images = document.getElementById("input-image").value

  if (project === "") {
    return alert("Please input Your Name Project")
  } else if (startDate === "") {
    return alert("Please Fill The Start Date")
  } else if (endDate === "") {
    return alert("lease Fill The End Date")
  } else if (description === "") {
    return alert("Pleasae Fill The Description")
  } else if (input.length === 0) {
    return alert("Please Choose The Technologies")
  } else if (images === "") {
    return alert("Please Choose The Image")
  }

  const nodeJsIcon = '<i class="bx bxl-nodejs"></i>'
  const reactIcon = '<i class="bx bxl-react"></i>'
  const javaIcon = '<i class="bx bxl-javascript"></i>'
  const typeIcon = '<i class="bx bxl-typescript"></i>'

  let nodejs = document.getElementById("nodejs").checked ? nodeJsIcon : ""
  let reactjs = document.getElementById("reactjs").checked ? reactIcon : ""
  let javascript = document.getElementById("javascript").checked ? javaIcon : ""
  let typescript = document.getElementById("typescript").checked ? typeIcon : ""

  image = URL.createObjectURL(image[0])

  let multiInput = document.querySelectorAll(".multi-input:checked")
  if (multiInput.length === 0) {
    return alert("Please Select At least One Technologies")
  }

  let start = new Date(startDate)
  let end = new Date(endDate)

  if (start > end) {
    return alert("You Fill End Date Before Start Date")
  }

  let difference = end.getTime() - start.getTime()
  let days = difference / (1000 * 3600 * 24)
  let weeks = Math.floor(days / 7)
  let months = Math.floor(weeks / 4)
  let years = Math.floor(months / 12)
  let duration = ""

  if (days > 0) {
    duration = days + " Hari"
  }
  if (weeks > 0) {
    duration = weeks + " Minggu"
  }
  if (months > 0) {
    duration = months + " Bulan"
  }
  if (years > 0) {
    duration = years + " Tahun"
  }

  let blog = {
    project,
    duration,
    description,
    nodejs,
    reactjs,
    javascript,
    typescript,
    image,
  }

  dataBlog.push(blog)
  renderBlog()
}

const renderBlog = () => {
  document.getElementById("contents").innerHTML = ""
  for (let i = 0; i < dataBlog.length; i++) {
    document.getElementById("contents").innerHTML += `
          <div class="d-flex flex-row flex-wrap ml-5">
          <div class="d-flex flex-row flex-wrap" style="margin-right: 10px;">
            <div class="card" style="width: 18rem">
              <img src="${dataBlog[i].image}" class="card-img-top" id="images" />
              <div class="card-body">
                <a href="blog-detail">
                  <h3 class="card-title text-left col-8 text-truncate">${dataBlog[i].project}</h3></a>
                <p>Durasi: ${dataBlog[i].duration}</p>
                <p class="card-text text-left">
                  ${dataBlog[i].description}
                </p>
                <div class="icon-project">
                  ${dataBlog[i].nodejs}
                  ${dataBlog[i].reactjs}
                  ${dataBlog[i].javascript}
                  ${dataBlog[i].typescript}
                </div>
                <div class="button-project">
                  <button class="btn-blue">Edit</button>
                  <button class="btn-blue">Delete</button>
                </div>
              </div>
            </div>
            </div>
          </div>`
  }
}
