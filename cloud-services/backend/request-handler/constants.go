package handler

const (
	htmlTemplate = `
    <!DOCTYPE html>
    <html>
    
    <head>
    
        <style>
            div.gallery {
                margin: 5px;
                border: 1px solid #ccc;
                float: left;
                width: 400px;
            }
    
            div.gallery:hover {
                border: 1px solid #777;
            }
    
            div.gallery img {
                width: 100%;
                height: auto;
            }
    
            div.desc {
                padding: 15px;
                text-align: center;
            }
        </style>
    
    </head>
    
    <body>
    
        <div class="gallery">
            <a target="_blank" href="https://images01.nicepage.com/a1389d7bc73adea1e1c1fb7e/77291429bf5052cc82b824bf/pexels-photo-1906818.jpeg">
                <img src="https://images01.nicepage.com/a1389d7bc73adea1e1c1fb7e/77291429bf5052cc82b824bf/pexels-photo-1906818.jpeg" alt="famous actor" width="1200" height="800">
            </a>
            <div class="desc">This is a picture from a famous actor</div>
            <hr>
            <div>
                Comments:
            </div>
            <hr>
            
            <!--cp1-->
    
            <form action="/newcomment1" method="post">
                <div>
                    <textarea name="comments" id="comments">Write you comment about this picture here</textarea>
                </div>
                <input type="submit" value="Submit">
            </form>
        </div>
    
        <div class="gallery">
            <a target="_blank" href="https://images01.nicepage.com/a1389d7bc73adea1e1c1fb7e/df4df7ec01155446b44e6fec/pexels-photo-9940878.jpeg">
                <img src="https://images01.nicepage.com/a1389d7bc73adea1e1c1fb7e/df4df7ec01155446b44e6fec/pexels-photo-9940878.jpeg" alt="professional CEO" width="1200" height="800">
            </a>
            <div class="desc">She is a professional CEO</div>
            <hr>
            <div>
                Comments:
            </div>
            <hr>
            
            <!--cp2-->
    
            <form action="/newcomment2" method="post">
                <div>
                    <textarea name="comments" id="comments">Write you comment about this picture here</textarea>
                </div>
                <input type="submit" value="Submit">
            </form>
        </div>
    
        <div class="gallery">
            <a target="_blank" href="https://images01.nicepage.com/a1389d7bc73adea1e1c1fb7e/3c6abf042f1b5649816bdd6c/pexels-photo-9640282.jpeg">
                <img src="https://images01.nicepage.com/a1389d7bc73adea1e1c1fb7e/3c6abf042f1b5649816bdd6c/pexels-photo-9640282.jpeg" alt="ugly moder" width="1200" height="800">
            </a>
            <div class="desc">She is an ugly model</div>
            
            <hr>
            <div>
                Comments:
            </div>
            <hr>
            
            <!--cp3-->
    
            <form action="/newcomment3" method="post">
                <div>
                    <textarea name="comments" id="comments">Write you comment about this picture here</textarea>
                </div>
                <input type="submit" value="Submit">
            </form>
        </div>
    
    
    </body>
    
    </html>`

	commentTemplate = `
<div>%s</div>
<audio controls>
    <source src="http://translate.google.com/translate_tts?ie=UTF-8&total=1&idx=0&textlen=32&client=tw-ob&q=%s&tl=en" type="audio/mpeg">
    Your browser does not support the audio element.
</audio>
<hr>
<!--cp%d-->`
)

var htmlTemplateVariable = htmlTemplate
