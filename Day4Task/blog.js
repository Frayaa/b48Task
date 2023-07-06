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
  let form = document
  

  if (project === "") {
    return alert("Please input Your Name Project")
  } else if (startDate === "") {
    return alert("Please fill the date")
  } else if (endDate === "") {
    return alert("please fill the date")
  } else if (description === "") {
    return alert("pleasae fill the description")
  } else if (input.length === 0) {
    return alert("please choose the technologies")
  } else if (images === "") {
    return alert("please choose the image")
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
  console.log(image)

  let multiInput = document.querySelectorAll(".multi-input:checked")
  if (multiInput.length === 0) {
    return alert("Select at least one technology used.")
  }

  let start = new Date(startDate)
  let end = new Date(endDate)

  if (start > end) {
    return alert("Your Fill End Date after Start Date")
  }

  let difference = end.getTime() - start.getTime()
  let days = difference / (1000 * 3600 * 24)
  let weeks = Math.floor(days / 7)
  let months = Math.floor(weeks / 4)
  let years = Math.floor(months / 12)
  let duration = ""

  if (days > 0) {
    duration = days + " hari"
  }
  if (weeks > 0) {
    duration = weeks + " minggu"
  }
  if (months > 0) {
    duration = months + " bulan"
  }
  if (years > 0) {
    duration = years + " tahun"
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
    <div class="container-project">
    <div class="container-card">
      <div class="image-project">
        <img src="${dataBlog[i].image}" />
      </div>
      <div class="title-project">
        <a href="blog-detail.html"><h3>${dataBlog[i].project}</h3></a>
        <p>Durasi: ${dataBlog[i].duration}</p>
      </div>
      <div class="content">
        <p> 
          ${dataBlog[i].description}
        </p>
      </div>
      <div class="icon-project">
        ${dataBlog[i].nodejs}
        ${dataBlog[i].reactjs}
        ${dataBlog[i].javascript}
        ${dataBlog[i].typescript}
      </div>
      <div class="button-project">
        <button>Edit</button>
        <button>Delete</button>
      </div>
      </div>
   </div>`
  }
}

const form = () => {
  document.getElementById('form').reset()
} 
