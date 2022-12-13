<script>
    import './assets/bootstrap.min.css'
    import './assets/bootstrap.bundle.min.js'
    import {IsOnline} from '../wailsjs/go/main/Scrap.js';

    let online  = true
    function scrap() {
	    IsOnline().then(result => {if (result === false)
		    {
			    online = false
			    console.log(online)
		    } else {
			    online = true
			    console.log(online)
		    }
	    })    
        
    }
     

</script>

<main>
	<h2>To Yuri with code!</h2> 
  	<button type="button" id="liveToastBtn" class="btn btn-outline-info" on:click={scrap} >Scrap</button>

	<div class="results"> 
		<p id="check">Found Internet</p>
		<p id="check">Scrapped website</p>
		<p id="check">Sent email</p>
	</div>

</main>

<style>
	button {
		margin-top: 3%;
		margin-left: 47%;
	}

	.results {
		margin-top: 8%;
		margin-left: 16%;
		width: 713px;
		height: 210px;
		border: 2px solid #57ACDC;
	}

	#check {
		color: #60C689;
		text-align: center;
		font-size: 23px;
	}
</style>
