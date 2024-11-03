<template>
    
        
        
        <div id="mouse-line" style="opacity: 0.99; position: fixed;">
            <vue-particles 
      color="#dedede"
      :particleOpacity="0.7"
      :particlesNumber="60"
      shapeType="circle"
      :particleSize="4"
      linesColor="#dedede"
      :linesWidth="1"
      :lineLinked="true"
      :lineOpacity="0.5"
      :linesDistance="150"
      :moveSpeed="3"
      :hoverEffect="true"
      hoverMode="grab"
      :clickEffect="true"
      clickMode="push"
    >
    
    </vue-particles>
        </div>
    <div class="container">
        <div style="opacity: 0.99; ">        
            <slot></slot>
        </div>

        </div>
</template>
    
<script>
export default {
    name: "BlankBG", 

    
}



</script>
    
<style scoped>
.container {
    margin-top: 10px;
}

#mouse-line{
    width: 100%;
    height: 100%

}


</style>

