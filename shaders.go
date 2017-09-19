package main

const (
	basicFragmentSRC = `
	#version 410
  precision mediump float;

	uniform sampler2D tex;
	uniform mat4 MVP;
	uniform mat4 MODEL;

	vec3 lightPos = vec3(0.0, 0.0, 0.0);
	const vec3 ambientColor = vec3(0.1, 0.1, 0.1)*3;
	const vec3 diffuseColor = vec3(0.1, 0.1, 0.1)*5;
	const vec3 specColor = vec3(0.1, 0.1, 0.1)*10;

	in vec3 normalInterp;
	in vec3 fragPos;
	in vec2 fragTexCoord;

  out vec4 finalColor;

	int mode = 1;

  void main() {
		vec3 normal = normalize(normalInterp);
	  vec3 lightDir = normalize(lightPos - fragPos);

	  float lambertian = max(dot(lightDir,normal), 0.0);
	  float specular = 0.0;

	  if(lambertian > 0.0) {

	    vec3 viewDir = normalize(-fragPos);

	    // this is blinn phong
	    vec3 halfDir = normalize(lightDir + viewDir);
	    float specAngle = max(dot(halfDir, normal), 0.0);
	    specular = pow(specAngle, 16.0);

	    // this is phong (for comparison)
	    if(mode == 2) {
	      vec3 reflectDir = reflect(-lightDir, normal);
	      specAngle = max(dot(reflectDir, viewDir), 0.0);
	      // note that the exponent is different here
	      specular = pow(specAngle, 4.0);
	    }
	  }

    vec4 surfaceColor = texture(tex, fragTexCoord);
		finalColor = vec4(ambientColor +
		                  lambertian * diffuseColor +
		                  specular * specColor, 1.0) * surfaceColor;

	  // finalColor = texture(tex, fragTexCoord);
    //finalColor = vec4(1,0,1,1);
  }` + "\x00"

	basicVertexSRC = `
	#version 410 core

	uniform mat4 MVP, MODEL, NormalMatrix;

	in vec3 vert;
	in vec2 vertTexCoord;
	in vec3 vertNormal;

	out vec3 fragPos;
	out vec2 fragTexCoord;
	out vec3 normalInterp;

	void main(){
	  vec4 fragPos4 = MODEL * vec4(vert, 1.0);
	  fragPos = vec3(fragPos4) / fragPos4.w;
		fragTexCoord = vertTexCoord;

		// mat4 normalMatrix = transpose(inverse(MODEL));
	  // normalInterp = vec3(normalMatrix * vec4(vertNormal, 0.0));
	  normalInterp = vec3(NormalMatrix * vec4(vertNormal, 0.0));

		gl_Position =  MVP * MODEL * vec4(vert, 1);
	}` + "\x00"
)
